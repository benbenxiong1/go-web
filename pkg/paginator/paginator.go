package paginator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-web/pkg/config"
	"go-web/pkg/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"strings"
)

// Paging 分页数据
type Paging struct {
	CurrentPage int // 当前页
	PerPage     int
	TotalPage   int
	TotalCount  int64
	NextPageUrl string
	PrevPageUrl string
}

// Paginator 分页操作类
type Paginator struct {
	BaseUrl    string
	PerPage    int
	Page       int
	Offset     int
	TotalCount int64
	TotalPage  int
	Sort       string
	Order      string

	query *gorm.DB
	ctx   *gin.Context
}

func Paginate(c *gin.Context, db *gorm.DB, data interface{}, baseUrl string, perPage int) Paging {
	// 初始化 Paginator 实例
	p := &Paginator{
		query: db,
		ctx:   c,
	}
	p.initProperties(perPage, baseUrl)

	// 查询数据库
	err := p.query.Preload(clause.Associations).
		Order(p.Sort + " " + p.Order).
		Limit(p.PerPage).
		Offset(p.Offset).
		Find(data).
		Error

	if err != nil {
		logger.LogIf(err)
		return Paging{}
	}

	return Paging{
		CurrentPage: p.Page,
		PerPage:     p.PerPage,
		TotalPage:   p.TotalPage,
		TotalCount:  p.TotalCount,
		NextPageUrl: p.getNextPageUrl(),
		PrevPageUrl: p.getPrevPageUrl(),
	}
}

// initProperties 初始化分页必须用到的属性，基于这些属性查询数据库
func (p *Paginator) initProperties(perPage int, baseUrl string) {
	p.BaseUrl = p.formatBaseUrl(baseUrl)
	p.PerPage = p.getPerPage(perPage)

	// 排序参数
	p.Order = p.ctx.DefaultQuery(config.Get("paging.url_query_order"), "asc")
	p.Sort = p.ctx.DefaultQuery(config.Get("paing.url_query_sort"), "id")

	p.TotalCount = p.getTotalCount()
	p.TotalPage = p.getTotalPage()
	p.Page = p.getCurrentPage()
	p.Offset = (p.Page - 1) * p.PerPage
}

func (p *Paginator) formatBaseUrl(url string) string {
	if strings.Contains(url, "?") {
		url = url + "&" + config.Get("paging.url_query_page") + "="
	} else {
		url = url + "?" + config.Get("paging.url_query_page") + "="
	}
	return url
}

func (p *Paginator) getPerPage(perPage int) int {
	// 优先使用 per_page 参数
	queryPerPage := p.ctx.Query(config.Get("paging.url_query_per_page"))
	if len(queryPerPage) > 0 {
		perPage = cast.ToInt(queryPerPage)
	}

	// 没有传参 使用默认
	if perPage <= 0 {
		perPage = config.GetInt("paging.perpage")
	}

	return perPage
}

// getTotalCount 返回的是数据库里的条数
func (p *Paginator) getTotalCount() int64 {
	var count int64
	if err := p.query.Count(&count).Error; err != nil {
		return 0
	}
	return count
}

// getTotalPage 计算总页数
func (p *Paginator) getTotalPage() int {
	if p.TotalCount == 0 {
		return 0
	}
	nums := int64(math.Ceil(float64(p.TotalCount) / float64(p.PerPage)))
	if nums == 0 {
		nums = 1
	}
	return int(nums)
}

func (p *Paginator) getCurrentPage() int {
	// 优先取用户请求的 page
	page := cast.ToInt(p.ctx.Query(config.Get("paging.url_query_page")))
	if page <= 0 {
		// 默认为1
		page = 1
	}

	// TotalPage 等于 0  意味着数据不够分页
	if p.TotalPage == 0 {
		return 0
	}

	// 请求页数大于总页数 返回总页数
	if page > p.TotalPage {
		return p.TotalPage
	}

	return page
}

// getPageLink 拼接分页链接
func (p *Paginator) getPageLink(page int) string {
	return fmt.Sprintf("%v%v&%s=%s&%s=%s&%s=%v",
		p.BaseUrl,
		page,
		config.Get("paging.url_query_sort"),
		p.Sort,
		config.Get("paging.url_query_order"),
		p.Order,
		config.Get("paging.url_query_per_page"),
		p.PerPage,
	)
}

// getPrevPageUrl 获取上一页的链接
func (p *Paginator) getPrevPageUrl() string {
	if p.Page <= 1 || p.Page > p.TotalPage {
		return ""
	}
	return p.getPageLink(p.Page - 1)
}

// getNextPageUrl 获取下一页链接
func (p *Paginator) getNextPageUrl() string {
	if p.TotalPage > p.Page {
		return p.getPageLink(p.Page + 1)
	}
	return ""
}
