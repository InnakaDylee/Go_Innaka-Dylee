package main


type user struct { // nama struct lebih baik menggunakan huruf kapital pada huruf pertama user

id int

username int  // beri tab size agar lebih rapih dan space antar line dihilangkan

password int

}


type userservice struct { // nama struct lebih baik menggunakan huruf kapital pada huruf pertama setiap kata userservice

t []user // beri tab size agar lebih rapih dan space antar line dihilangkan dan nama field dijelaskan lagi lebih detail seperti attribut

}


func (u userservice) getallusers() []user { // nama fungsi lebih baik menggunakan huruf kapital pada huruf pertama agar dapat digunakan diluar folder/direktori setiap kata getallusers
	  										// nama parameter u dijelaskan lebih detail seperti user
return u.t // beri tab size agar lebih rapih dan space antar line dihilangkan dan nama field t dijelaskan lebih detail seperti attribut

}


func (u userservice) getuserbyid(id int) user { // nama fungsi lebih baik menggunakan huruf kapital pada huruf pertama agar dapat digunakan diluar folder/direktori setiap kata getuserbyid
												// nama parameter u dijelaskan lebih detail seperti user
for _, r := range u.t { // beri tab size agar lebih rapih dan space antar line dihilangkan dan nama variable t dijelaskan lebih detail seperti value

if id == r.id { // beri tab size agar lebih rapih dan space antar line dihilangkan dan nama variable id dijelaskan lebih detail seperti userId

return r // beri tab size agar lebih rapih dan space antar line dihilangkan

}

}


return user{}

}