// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"fmt"
	"log"
	"math/rand"

	"baliance.com/gooxml/chart"
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/spreadsheet"

	sml "baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
)

func main() {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	// Create all of our data
	row := sheet.AddRow()
	/*	hdrStyle := ss.StyleSheet.AddCellStyle()
		pf := ss.StyleSheet.Fills().AddPatternFill()
		pf.SetFgColor(color.LightGray)
		hdrStyle.SetFill(pf)
	*/
	row.AddCell().SetString("Item")
	row.AddCell().SetString("Price")
	row.AddCell().SetString("# Sold")
	row.AddCell().SetString("Total")
	for r := 0; r < 5; r++ {
		row := sheet.AddRow()
		row.AddCell().SetString(fmt.Sprintf("Product %d", r+1))
		row.AddCell().SetNumber(float64(rand.Intn(50)) / 10.0)
		row.AddCell().SetNumber(float64(rand.Intn(50) + 1))
		row.AddCell().SetFormulaRaw(fmt.Sprintf("C%d*B%d", r+2, r+2))
	}

	// add an auto-filter
	sheet.SetAutoFilter("A1:D6")

	// conditional formatting
	// total column
	cf := sheet.AddConditionalFormatting([]string{"D2:D6"})
	rule := cf.AddRule()
	db := rule.SetDataBar()
	db.AddFormatValue(sml.ST_CfvoTypeMin, "0")
	db.AddFormatValue(sml.ST_CfvoTypeMax, "0")
	db.SetColor(color.Blue)

	// sold column
	cf = sheet.AddConditionalFormatting([]string{"B2:B6"})
	rule = cf.AddRule()
	cs := rule.SetColorScale()
	cs.AddFormatValue(sml.ST_CfvoTypeMin, "0")
	cs.AddFormatValue(sml.ST_CfvoTypePercentile, "50")
	cs.AddFormatValue(sml.ST_CfvoTypeMax, "0")
	cs.AddGradientStop(color.SuccessGreen)
	cs.AddGradientStop(color.Orange)
	cs.AddGradientStop(color.Red)

	// Charts need to reside in a drawing
	dwng := ss.AddDrawing()
	chrt1, anc1 := dwng.AddChart()
	chrt2, anc2 := dwng.AddChart()
	addBar3DChart(chrt1)
	addLineChart(chrt2)
	anc1.SetWidth(9)
	anc1.MoveTo(5, 1)
	anc2.MoveTo(1, 23)

	// and finally add the chart to the sheet
	sheet.SetDrawing(dwng)

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}
	ss.SaveToFile("complex.xlsx")
}

func addBar3DChart(chrt chart.Chart) {
	chrt.AddTitle().SetText("Bar Chart")
	lc := chrt.AddBar3DChart()
	priceSeries := lc.AddSeries()
	priceSeries.SetText("Price")
	// Set a category axis reference on the first series to pull the product names
	priceSeries.CategoryAxis().SetLabelReference(`'Sheet 1'!A2:A6`)
	priceSeries.Values().SetReference(`'Sheet 1'!B2:B6`)

	soldSeries := lc.AddSeries()
	soldSeries.SetText("Sold")
	soldSeries.Values().SetReference(`'Sheet 1'!C2:C6`)

	totalSeries := lc.AddSeries()
	totalSeries.SetText("Total")
	totalSeries.Values().SetReference(`'Sheet 1'!D2:D6`)

	// the line chart accepts up to two axes
	ca := chrt.AddCategoryAxis()
	va := chrt.AddValueAxis()
	lc.AddAxis(ca)
	lc.AddAxis(va)

	ca.SetCrosses(va)
	va.SetCrosses(ca)
}

func addLineChart(chrt chart.Chart) {
	chrt.AddTitle().SetText("Line Chart")
	lc := chrt.AddLine3DChart()
	priceSeries := lc.AddSeries()
	priceSeries.SetText("Price")
	// Set a category axis reference on the first series to pull the product names
	priceSeries.CategoryAxis().SetLabelReference(`'Sheet 1'!A2:A6`)
	priceSeries.Values().SetReference(`'Sheet 1'!B2:B6`)

	soldSeries := lc.AddSeries()
	soldSeries.SetText("Sold")
	soldSeries.Values().SetReference(`'Sheet 1'!C2:C6`)

	totalSeries := lc.AddSeries()
	totalSeries.SetText("Total")
	totalSeries.Values().SetReference(`'Sheet 1'!D2:D6`)

	// the line chart accepts up to two axes
	ca := chrt.AddCategoryAxis()
	va := chrt.AddValueAxis()
	lc.AddAxis(ca)
	lc.AddAxis(va)

	ca.SetCrosses(va)
	va.SetCrosses(ca)
}