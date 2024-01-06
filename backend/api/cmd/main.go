package main

import (
	"log"

	"github.com/brnocorreia/ufba-enhacer/internal/configs/database"
	"github.com/brnocorreia/ufba-enhacer/internal/controller"
	"github.com/brnocorreia/ufba-enhacer/internal/controller/routes"
	"github.com/brnocorreia/ufba-enhacer/internal/model/repository"
	"github.com/brnocorreia/ufba-enhacer/internal/model/repository/entity"
	"github.com/brnocorreia/ufba-enhacer/internal/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var models = []interface{}{&entity.Course{}, &entity.Discipline{}, &entity.Prerequisite{}, &entity.CoursesDiscipline{}, &entity.User{}}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.NewDBConnection()
	if err != nil {
		log.Fatal("Error connecting to db")
	}

	db.AutoMigrate(models...)

	usersRepository := repository.NewUserRepository(db)
	userService := service.NewUserDomainService(usersRepository)
	userController := controller.NewUserControllerInterface(userService)

	router := gin.Default()

	routes.InitUserRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	// user := entity.User{
	// 	Model: gorm.Model{},
	// 	Name:  "Bruno Correia",
	// 	Email: "bruninhocorreia2013@gmail.com",
	// }

	// db.Create(&user)

	// fmt.Println(user.ID)

	// new_course := entity.Course{
	// 	Code:        3160,
	// 	Name:        "ADMINISTRAÇÃO",
	// 	Shift:       "Diurno",
	// 	MinDuration: 4.0,
	// 	MaxDuration: 8.0,
	// 	LegalBase:   "DATA DE INÍCIO: 03.03.1961. RECONHECIMENTO: PARECER CFE Nº 167 DE 27.12.1962. DIRETRIZES CURRICULARES: PARECER CNE Nº 04 de 13.07.2005",
	// 	Description: "Ao Administrador compete o plajejamento e a implantação de estruturas ognizacionais, dos métodos de trabalho e da utilização dos recursos humanos, materiais e financeiros disponíveis, com vistas à consecução dos objetivos da organicação (governamental ou privada) a que serve. Para isso, lança mão e desenvolve técnicas de administração de pessoal, de liderança e treinamento, de interrelacionamento de pessoas e grupos, de administração financeira e contábil, programação, orçamentação, pesquisa, processamento de dados, controle de produção, estudo de mercado, vendas, etc.",
	// 	ObWorkload:  2805,
	// 	OpWorkload:  240,
	// 	AcWorkload:  100,
	// 	Disciplines: []entity.Discipline{},
	// }
}
