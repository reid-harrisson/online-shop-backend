package utils

import (
	"fmt"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"

	"github.com/brianvoe/gofakeit/v6"
)

type Fruit struct {
	Name        string  `fake:"{fruit}"`
	Description string  `fake:"{loremipsumsentence:10}"`
	Price       float64 `fake:"{price:1,10}"`
}

func generateFruit() []string {
	var f Fruit
	gofakeit.Struct(&f)

	froot := []string{}
	froot = append(froot, f.Name)
	froot = append(froot, f.Description)
	froot = append(froot, fmt.Sprintf("%.2f", f.Price))

	return froot
}

func FruitList(length int) [][]string {
	var fruits [][]string

	for i := 0; i < length; i++ {
		ff := generateFruit()
		fruits = append(fruits, ff)
	}

	return fruits
}

func BuildHeading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				// err := m.FileImage("images/logo_div_rhino.jpg", props.Rect{
				// 	Center:  true,
				// 	Percent: 75,
				// })

				// if err != nil {
				// 	fmt.Println("Image file was not loaded 😱 - ", err)
				// }
			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Prepared for you by the Div Rhino Fruit Company", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
				Color: getDarkPurpleColor(),
			})
		})
	})
}

func BuildFruitList(m pdf.Maroto) {
	tableHeadings := []string{"Fruit", "Description", "Price"}
	contents := FruitList(20)
	lightPurpleColor := getLightPurpleColor()

	m.SetBackgroundColor(getTealColor())
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Products", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 7, 2},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3, 7, 2},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 false,
	})

}

func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}

func getLightPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}

func getTealColor() color.Color {
	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}
