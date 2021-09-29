package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/techagentn/cto/model"
	"html/template"
	"log"
	"net/http"
)

type Post struct {
	userId string
	Title string
	Body string
}

var templateFiles = []string{
	 "./ui/html/home.page.gohtml",
	 "./ui/html/base.layout.gohtml",
	"./ui/html/footer.partial.gohtml",
	//"./ui/html/main.page.gohtml.gohtml",

}

func (app *application) home(w http.ResponseWriter, r * http.Request){
	if r.URL.Path != "/" {
		//http.NotFound(w, r)
		app.notFound(w)
		return
	}
	ts, err := template.ParseFiles(templateFiles...)
	if err != nil {
		//app.errLog.Panicln(err.Error())
		//http.Error(w, "internal server error", http.StatusInternalServerError)
		app.serverError(w, err) // using custom helpers
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		//app.errLog.Panicln(err.Error())
		//http.Error(w, "internal server error", http.StatusInternalServerError)
		app.serverError(w, err) // using custom helpers for server error
		return
	}
}


func (b *Post) GetAllContents()(blogPosts  []Post, err error)   {
	rows, err := model.Db.Query(`SELECT userId,Title,Body FROM tech`)
	if err != nil{
		return
	}
	for rows.Next(){
		bp := Post{}
		err = rows.Scan(&bp.userId,&bp.Title,&bp.Body)
		if err != nil{
			return
		}
		blogPosts = append(blogPosts,bp)
	}
	rows.Close()
	return

}



func (app *application) blog(w http.ResponseWriter, r * http.Request){
	templateFiles := []string{
		"./ui/html/blog.page.gohtml",
		"./ui/html/base.layout.gohtml",
		"./ui/html/footer.partial.gohtml",
		//"./ui/html/main.page.gohtml.gohtml",
	}
	ts, err := template.ParseFiles(templateFiles...)
	if err != nil {
		//app.errLog.Panicln(err.Error())
		//http.Error(w, "internal server error", http.StatusInternalServerError)
		app.serverError(w, err) // using custom helpers
		return
	}
	blogPost := Post{}
	users, err := blogPost.GetAllContents()
	if err != nil{
		app.errLog.Println("Can't read from the database")
	}
	err = ts.Execute(w, users)
}
func (b *Post) CreateBlog()  {
	stmt, err := model.Db.Prepare(`INSERT INTO tech(userId,Title, Body) VALUES($1, $2, $3)`)
	defer stmt.Close()
	if err != nil{
		log.Panicln(err.Error())
	}
	_, err = stmt.Exec(b.userId, b.Title, b.Body)
	if err != nil {
		panic(err.Error())
		//http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
}
func (app *application) processPForm(w http.ResponseWriter, r * http.Request){

	//templateFiles := []string{
	//	"./ui/html/blog.page.gohtml",
	//	"./ui/html/base.layout.gohtml",
	//	"./ui/html/footer.partial.gohtml",
	//}
	//ts, err := template.ParseFiles(templateFiles...)
	//if err != nil {
	//	//app.errLog.Panicln(err.Error())
	//	//http.Error(w, "internal server error", http.StatusInternalServerError)
	//	app.serverError(w, err) // using custom helpers
	//	return
	//}
// Database store
	r.ParseForm()
	//Store data
	blogPost := &Post{
		userId: uuid.New().String(),
		Title: r.Form.Get("title"),
		Body: r.Form.Get("body"),
	}
	//myPost = append(myPost, blogPost)
	//Insert values
	blogPost.CreateBlog()
	//err = ts.Execute(w, blogPost)
	//fmt.Println(err)
	http.Redirect(w, r, "/blog", http.StatusFound)
}

func (b *Post) Delete(val string) (err error) {
	_, err = model.Db.Query(`delete from tech where userId = $1`,val)
	return err
}

func (b *Post) Del(val string) (err error) {
	_, err = model.Db.Query(`delete from tech where userId = $1`,val)
	return err
}

func (app *application) Delete(w http.ResponseWriter, r *http.Request)  {
	blog := Post{}
	id := chi.URLParam(r,"userId")

	err := blog.Del(id)
	app.infoLog.Printf(id)

	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w,r,"/blog",http.StatusFound)
	//71697cc7-c919-4d0d-b3ad-97f6fc6b51b2


}