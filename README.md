# LoginApp with go and vuejs

step 1
go mod init backend/webapp.go

step2
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/joho/godotenv

step3
go run webapp.go
this opens the commnd prompt to run
/api/register

{
"username":"user",
"password":"pass"
}

/api/login

{
"username":"user",
"password":"pass"
}

# for the frontend

npm install -g @vue/cli

vue create vue-frontend
npm install axios
vue add component Login
# login-app-2
