# Summary

## Framework Echo Golang
* Golang memiliki sejumlah perpustakaan pihak ketiga yang berguna, termasuk:

    * Echo
    * Go Kit
    * Logrus
    * gorm.io
    * Cobra
* Echo adalah suatu kerangka kerja web yang ringan dan sangat efisien yang digunakan untuk membangun aplikasi web dengan bahasa pemrograman Go. Kerangka kerja ini dirancang untuk mencapai kecepatan dan kemanfaatan maksimal, dan seringkali digunakan untuk membuat aplikasi web berbasis RESTful.

* Ada beberapa keuntungan yang diberikan oleh Echo:

    * Router yang telah dioptimalkan.
    * Kemampuan merender data.
    * Data binding yang mudah.
    * Skalabilitas.
    * Kemampuan untuk menggunakan middleware.

* Contoh penggunaan dasar untuk routing dan pengendali (controller) dalam Echo:

    ```
    go
    Copy code
    import (
    "net/http"
    "github.com/labstack/echo"
    )

    func main() {
    e := echo.New()
    e.GET("/user/:id/:age", UserController)
    e.POST("/user", RegisterController)
    e.Start(":8000")
    }

    func UserController(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]interface{}{
        "user": user,
        "search": search,
        "short": short,
    })
    }
    ```

* Echo juga memudahkan penggunaan data dengan melakukan rendering sebagai berikut:

    ```
    go
    Copy code
    import (
    "net/http"
    "strconv"
    "github.com/labstack/echo"
    )

    type User struct {
    Id    int    `json:"id" form:"id"`
    Age   int    `json:"age" form:"age"`
    Email string `json:"email" form:"email"`
    Name  string `json:"name" form:"name"`
    }

    func main() {
    e := echo.New()
    e.GET("/user/:id/:age", UserController)
    e.POST("/user", RegisterController)
    e.Start(":8000")
    }

    func UserController(c echo.Context) error {
    user := User{Id: id, Age: age, Email: "anggikiyowo@gmail.com", Name: "Anggi Cantik"}
    return c.JSON(http.StatusOK, map[string]interface{}{
        "user": user,
        "search": search,
        "short": short,
    })
    }
    ```

* Untuk mengambil data dari permintaan, Anda bisa melakukannya seperti berikut:
    ```
    go
    Copy code
    type User struct {
    Id    int    `json:"id" form:"id"`
    Age   int    `json:"age" form:"age"`
    Email string `json:"email" form:"email"`
    Name  string `json:"name" form:"name"`
    }

    func main() {
    e := echo.New()
    e.GET("/user/:id/:age", UserController)
    e.POST("/user", RegisterController)
    e.Start(":8000")
    }

    func UserController(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    age, _ := strconv.Atoi(c.Param("age"))
    user := User{Id: id, Age: age, Email: "anggikiyowo@gmail.com", Name: "Anggi Cantik"}
    return c.JSON(http.StatusOK, map[string]interface{}{
        "user": user,
        "search": search,
        "short": short,
    })
    }
    ```

* Pengikatan (binding) data juga bisa dicapai dengan mudah, seperti yang ditunjukkan dalam fungsi RegisterController:
    ```
    go
    Copy code
    func RegisterController(c echo.Context) error {
    user := User{}
    c.Bind(&user)
    return c.JSON(http.StatusOK, map[string]interface{}{
        "user": user,
    })
    }
    ```