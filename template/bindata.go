package template

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetse7aeeb3a3cdab288001d1040d591acc0d164d7a4 = "    GetBy{{.FieldSearch}}(ctx context.Context,limit,page int,param{{.FieldSearch}} string) (interface{}, error)"
var _Assetsbcaadb4805cbd0f1a726bb64657ce83075b6309f = "package {{.Package}}\n\nimport (\n\t\"context\"\n    \"math\"\n    \"strconv\"\n)\n\n/********************************************************************************\n* Temancode {{.Name}} Service Package                                             *\n*                                                                               *\n* Version: 1.0.0                                                                *\n* Date:    2023-01-05                                                           *\n* Author:  Waluyo Ade Prasetio                                                  *\n* Github:  https://github.com/abdullahPrasetio                                  *\n********************************************************************************/\n\ntype service struct {\n\trepository Repository\n}\ntype Service interface {\n\tGetAll(ctx context.Context,limit,page int) (interface{}, error)\n\t// Add Function interface\n    // End Function interface\n}\n\nfunc NewService(repo Repository) *service {\n\treturn &service{repository: repo}\n}\n\ntype Paginate struct {\n\tData         interface{}              `json:\"data\"`\n\tTotalPage    int                      `json:\"total_page\"`\n\tCurrentPage  int                      `json:\"current_page\"`\n\tStartPage    int                      `json:\"start_page\"`\n\tLastPage     int                      `json:\"last_page\"`\n\tLimit        int                      `json:\"limit\"`\n\tNextPage     string                   `json:\"next_page\"`\n\tPreviousPage string                   `json:\"previous_page\"`\n\tLinks        []map[string]interface{} `json:\"links\"`\n}\n\nfunc (s *service) GetAll(ctx context.Context,limit,page int) (interface{}, error) {\n\tvar err error\n\tvar result []{{.Name}}\n\tvar total int\n\t// menentukan offset\n    offset := (page - 1) * limit\n\tresult,total, err = s.repository.Get(ctx,ParamPagination{uint64(limit), uint64(offset)})\n\tif err != nil {\n\t\treturn result, err\n\t}\n\tif limit>0{\n        paginate := CreatePagination(result, limit, page, total)\n\n        return paginate, nil\n    }\n    return result, nil\n}\n\nfunc CreatePagination(data interface{}, limit, page, total int) Paginate {\n\t//change to hostname\n\turl := \"http://localhost:8025/api/v1/{{.NameTable}}\"\n\tinterval := 1\n\tcurrentPage := page\n\ttotalPages := int(math.Ceil(float64(total) / float64(limit)))\n\tlastPage := totalPages\n\tstartPage := 1\n\tnextPageUrl := \"\"\n\tnextPage := strconv.Itoa(currentPage + 1)\n\tprevPageUrl := \"\"\n\tprevPage := strconv.Itoa(currentPage - 1)\n\n\tlinks := []map[string]interface{}{}\n\n\tstart := currentPage - interval\n\tif start < startPage {\n\t\tstart = startPage\n\t}\n\n\tif startPage == currentPage {\n\t\tprevPageUrl = \"\"\n\t\tprevPage = strconv.Itoa(currentPage)\n\t} else {\n\t\tprevPageUrl = url + \"?page=\" + prevPage\n\t\tlink := map[string]interface{}{\n\t\t\t\"url\":           prevPageUrl,\n\t\t\t\"isCurrentPage\": false,\n\t\t\t\"page\":          prevPage,\n\t\t\t\"label\":         \"Prev\",\n\t\t}\n\t\tlinks = append(links, link)\n\t}\n\tend := currentPage + interval\n\tif end > lastPage {\n\t\tend = lastPage\n\t}\n\n\tfor i := start; i <= end; i++ {\n\t\tisCurrentPage := false\n\t\tif i == page {\n\t\t\tisCurrentPage = true\n\t\t}\n\t\tlink := map[string]interface{}{\n\t\t\t\"url\":           url + \"?page=\" + strconv.Itoa(i),\n\t\t\t\"isCurrentPage\": isCurrentPage,\n\t\t\t\"page\":          strconv.Itoa(i),\n\t\t\t\"label\":         strconv.Itoa(i),\n\t\t}\n\t\tlinks = append(links, link)\n\t}\n\n\tif lastPage == currentPage {\n\t\tnextPageUrl = \"\"\n\t\tnextPage = strconv.Itoa(currentPage)\n\t} else {\n\t\tnextPageUrl = url + \"?page=\" + nextPage\n\t\tlink := map[string]interface{}{\n\t\t\t\"url\":           nextPageUrl,\n\t\t\t\"isCurrentPage\": false,\n\t\t\t\"page\":          nextPage,\n\t\t\t\"label\":         \"Next\",\n\t\t}\n\t\tlinks = append(links, link)\n\t}\n\tpaginate := Paginate{\n\t\tData:         data,\n\t\tTotalPage:    totalPages,\n\t\tLimit:        limit,\n\t\tCurrentPage:  page,\n\t\tStartPage:    startPage,\n\t\tLastPage:     lastPage,\n\t\tNextPage:     nextPage,\n\t\tPreviousPage: prevPage,\n\t\tLinks:        links,\n\t}\n\treturn paginate\n}\n\n\n// Add New Function\n// End New Function\n"
var _Assetsca066eae1aa9527c147d55aac89a5b60a034239f = "func (s *service) GetBy{{.FieldSearch}}(ctx context.Context,limit,page int,param{{.FieldSearch}} string) (interface{}, error) {\n\tvar err error\n\tvar result []{{.Name}}\n\tvar total int\n\t// menentukan offset\n    offset := (page - 1) * limit\n\tresult,total, err = s.repository.GetBy{{.FieldSearch}}(ctx,ParamPagination{uint64(limit), uint64(offset)},param{{.FieldSearch}})\n\tif err != nil {\n\t\treturn result, err\n\t}\n\tif limit>0{\n        paginate := CreatePagination(result, limit, page, total)\n\n        return paginate, nil\n    }\n    return result, nil\n}"
var _Assets3eaedf941475e541b28b0aa28c3bf6236a3d8df3 = "package {{.Package}}\r\n\r\nimport \"time\"\r\n\r\n/********************************************************************************\r\n* Temancode {{.Name}} Entity Package                                              *\r\n*                                                                               *\r\n* Version: 1.0.0                                                                *\r\n* Date:    2023-01-05                                                           *\r\n* Author:  Waluyo Ade Prasetio                                                  *\r\n* Github:  https://github.com/abdullahPrasetio                                  *\r\n********************************************************************************/\r\n\r\n// Please Do not edit this file if you don't need\r\n// Timeout query in seconds\r\nvar timeout time.Duration = 20\r\n// Entity start\r\n\r\n// Entity end\r\n"
var _Assetsb3b087e47f32ebc17aa0060b3fbf4c0d11697cb1 = "package {{.Package}}\r\n\r\n\r\n/********************************************************************************\r\n* Temancode {{.Name}} Repository Package                                          *\r\n*                                                                               *\r\n* Version: 1.0.0                                                                *\r\n* Date:    2023-01-05                                                           *\r\n* Author:  Waluyo Ade Prasetio                                                  *\r\n* Github:  https://github.com/abdullahPrasetio                                  *\r\n********************************************************************************/\r\n\r\n\r\nimport (\r\n\t\"context\"\r\n\t\"database/sql\"\r\n\t\"fmt\"\r\n\tsq \"github.com/Masterminds/squirrel\"\r\n\t\"time\"\r\n)\r\n// Template created by Waluyo Ade Prasetio@Temancode\r\n\r\nvar (\r\n\tSQL         string\r\n\tqueryInsert sq.InsertBuilder\r\n\tquerySelect sq.SelectBuilder\r\n\tqueryCount  sq.SelectBuilder\r\n\terr         error\r\n\targuments   []interface{}\r\n)\r\n\r\ntype repository struct {\r\n\tdb    *sql.DB\r\n}\r\n\r\ntype Repository interface {\r\n\tGet(ctx context.Context, paginate ParamPagination) ([]{{.Name}},int, error)\r\n\tCount(ctx context.Context) (int, error)\r\n    // Add Function interface\r\n    // End Function interface\r\n}\r\n\r\ntype ParamPagination struct {\r\n\tLimit  uint64\r\n\tOffset uint64\r\n}\r\n\r\nfunc NewRepository(db *sql.DB) *repository {\r\n\treturn &repository{db: db}\r\n}\r\n\r\nfunc (r repository) Get(ctx context.Context, paginate ParamPagination) ([]{{.Name}},int, error) {\r\n\tvar total int\r\n\tresults := []{{.Name}}{}\r\n\tnewctx, cancel := context.WithTimeout(ctx, timeout*time.Second)\r\n\tdefer cancel()\r\n\tfieldQuery := {{.FieldQuery}}\r\n    querySelect = sq.Select(fieldQuery...).From(\"{{.NameTable}}\")\r\n\tif paginate.Limit != 0 {\r\n\t\tquerySelect = querySelect.Limit(paginate.Limit).Offset(paginate.Offset)\r\n\t\tqueryCount = sq.Select(\"COUNT(*)\").From(\"users\")\r\n        sqlCountQuery, args, err := queryCount.ToSql()\r\n        if err != nil {\r\n            panic(err.Error())\r\n        }\r\n        err = r.db.QueryRowContext(newctx, sqlCountQuery, args...).Scan(&total)\r\n        if err != nil {\r\n            return results, total, err\r\n        }\r\n\t}\r\n\tSQL, arguments, err = querySelect.ToSql()\r\n\tfmt.Println(SQL)\r\n\tif err != nil {\r\n\t\tpanic(err.Error())\r\n\t}\r\n\trows, err := r.db.QueryContext(newctx, SQL, arguments...)\r\n\tif err != nil {\r\n\t\treturn results,total, err\r\n\t}\r\n\tdefer rows.Close()\r\n\tfor rows.Next() {\r\n\t\tresult := {{.Name}}{}\r\n\t\terr := rows.Scan({{.FieldScan}})\r\n\t\tif err != nil {\r\n\t\t\treturn results,total, err\r\n\t\t}\r\n\t\tresults = append(results, result)\r\n\t}\r\n\r\n\treturn results,total, nil\r\n}\r\n\r\nfunc (r repository) Count(ctx context.Context) (int, error) {\r\n\tvar result int\r\n\tnewctx, cancelFunc := context.WithTimeout(ctx, timeout*time.Second)\r\n\tdefer cancelFunc()\r\n\tquerySelect = sq.Select(\"COUNT(*)\").From(\"{{.NameTable}}\")\r\n\tSQL, arguments, err = querySelect.ToSql()\r\n\tif err != nil {\r\n\t\tpanic(err.Error())\r\n\t}\r\n\terr = r.db.QueryRowContext(newctx, SQL).Scan(&result)\r\n\tif err != nil {\r\n\t\treturn result, err\r\n\t}\r\n\treturn result, nil\r\n}\r\n\r\n// Add New Function\r\n// End New Function\r\n"
var _Assets3de0d987c70c43333956ad79bb3d55e1cb9045a5 = "func (r repository) GetBy{{.FieldSearch}}(ctx context.Context, paginate ParamPagination, param{{.FieldSearch}} string) ([]{{.Name}},int, error) {\n\tvar total int\n\tresults := []{{.Name}}{}\n\n\tnewctx, cancel := context.WithTimeout(ctx, timeout*time.Second)\n\tdefer cancel()\n\tfieldQuery := {{.FieldQuery}}\n\tquerySelect = sq.Select(fieldQuery...).From(\"{{.NameTable}}\").Where(sq.Eq{\"{{.FieldSearchDatabase}}\": param{{.FieldSearch}}})\n\tif paginate.Limit != 0 {\n\t\tquerySelect = querySelect.Limit(paginate.Limit).Offset(paginate.Offset)\n\t\tqueryCount = sq.Select(\"COUNT(*)\").From(\"{{.NameTable}}\").Where(sq.Eq{\"{{.FieldSearchDatabase}}\": param{{.FieldSearch}}})\n        sqlCountQuery, args, err := queryCount.ToSql()\n        if err != nil {\n            panic(err.Error())\n        }\n        err = r.db.QueryRowContext(newctx, sqlCountQuery, args...).Scan(&total)\n        if err != nil {\n            return results, total, err\n        }\n\t}\n\tSQL, arguments, err = querySelect.ToSql()\n\tfmt.Println(param{{.FieldSearch}}, arguments)\n\tfmt.Println(SQL)\n\tif err != nil {\n\t\tpanic(err.Error())\n\t}\n\trows, err := r.db.QueryContext(newctx, SQL, arguments...)\n\tif err != nil {\n\t\treturn results, total, err\n\t}\n\tdefer rows.Close()\n\tfor rows.Next() {\n\t\tresult := {{.Name}}{}\n\t\terr := rows.Scan({{.FieldScan}})\n\t\tif err != nil {\n\t\t\treturn results, total, err\n\t\t}\n\t\tresults = append(results, result)\n\t}\n\n\treturn results, total, nil\n}"
var _Assets66ceaf54548ef449ce3f08fbc16f596cbfc08166 = "    GetBy{{.FieldSearch}}(ctx context.Context, paginate ParamPagination, param{{.FieldSearch}} string) ([]{{.Name}},int, error)"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"template"}, "/template": []string{"entity.txt", "repository.txt", "repository_search.txt", "repository_search_interface.txt", "service.txt", "service_search.txt", "service_search_interface.txt"}}, map[string]*assets.File{
	"/template/repository.txt": &assets.File{
		Path:     "/template/repository.txt",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1676216198, 1676216198284659400),
		Data:     []byte(_Assetsb3b087e47f32ebc17aa0060b3fbf4c0d11697cb1),
	}, "/template/repository_search.txt": &assets.File{
		Path:     "/template/repository_search.txt",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1676216410, 1676216410420365100),
		Data:     []byte(_Assets3de0d987c70c43333956ad79bb3d55e1cb9045a5),
	}, "/template/repository_search_interface.txt": &assets.File{
		Path:     "/template/repository_search_interface.txt",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1676201983, 1676201983718328300),
		Data:     []byte(_Assets66ceaf54548ef449ce3f08fbc16f596cbfc08166),
	}, "/template/service.txt": &assets.File{
		Path:     "/template/service.txt",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1676253791, 1676253791187949100),
		Data:     []byte(_Assetsbcaadb4805cbd0f1a726bb64657ce83075b6309f),
	}, "/template/service_search.txt": &assets.File{
		Path:     "/template/service_search.txt",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1676216271, 1676216271634121800),
		Data:     []byte(_Assetsca066eae1aa9527c147d55aac89a5b60a034239f),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1676254833, 1676254833911758900),
		Data:     nil,
	}, "/template/entity.txt": &assets.File{
		Path:     "/template/entity.txt",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1676216663, 1676216663032683200),
		Data:     []byte(_Assets3eaedf941475e541b28b0aa28c3bf6236a3d8df3),
	}, "/template": &assets.File{
		Path:     "/template",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1676216931, 1676216931672005200),
		Data:     nil,
	}, "/template/service_search_interface.txt": &assets.File{
		Path:     "/template/service_search_interface.txt",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1676215422, 1676215422840461200),
		Data:     []byte(_Assetse7aeeb3a3cdab288001d1040d591acc0d164d7a4),
	}}, "")