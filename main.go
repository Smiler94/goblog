package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    //w.Header().Set("Content-type", "text/html;charset=utf-8")
    fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog！</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    //w.Header().Set("Content-type", "text/html;charset=utf-8")
    fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
            "<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    //w.Header().Set("Content-type", "text/html;charset=utf-8")
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}

func articlesShowsHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    fmt.Fprint(w, "文章 ID：" + id)
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "访问文章列表")
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "创建新的文章")
}

func forceHtmlMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 1、设置标头
        w.Header().Set("Content-type", "text/html; charset=utf-8")
        // 2、继续处理请求
        next.ServeHTTP(w, r)
    })
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/home", homeHandler).Methods("GET").Name("home")
    router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")

    router.HandleFunc("/articles/{id:[0-9]+}", articlesShowsHandler).Methods("GET").Name("articles.show")
    router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("atciles.index")
    router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")

    // 自定义 404页面
    router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

    // 中间件，强制内容类型为 HTML
    router.Use(forceHtmlMiddleware)

    // 通过命名路由获取URL示例
    homeUrl, _ := router.Get("home").URL()
    fmt.Println("homeURL:", homeUrl)
    articleUrl, _ := router.Get("articles.show").URL("id", "23")
    fmt.Println("articleURL:", articleUrl)

    http.ListenAndServe(":8080", router)
}