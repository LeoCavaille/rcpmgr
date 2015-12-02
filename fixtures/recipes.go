package fixtures

import "github.com/LeoCavaille/rcpmgr/model"

var GaletteDesRois model.Recipe = model.Recipe{
	Title:    "Galette des Rois",
	Author:   "leo@cavaille.net",
	Category: "Desserts",
	Ingredients: []model.Ingredient{
		model.Ingredient{
			Quantity: 125,
			Unit:     "gram",
			Name:     "butter",
			Category: "dairy",
		},
		model.Ingredient{
			Quantity: 220,
			Unit:     "gram",
			Name:     "almond flour",
		},
		model.Ingredient{
			Quantity: 150,
			Unit:     "gram",
			Name:     "sugar",
		},
		model.Ingredient{
			Quantity: 2,
			Unit:     "tablespoon",
			Name:     "crème fraîche",
			Category: "dairy",
		},
		model.Ingredient{
			Quantity: 2,
			Name:     "whole egg",
			Category: "dairy",
		},
		model.Ingredient{
			Quantity: 1,
			Name:     "egg yolk",
			Category: "dairy",
		},
		model.Ingredient{
			Quantity: 1,
			Unit:     "teaspoon",
			Name:     "rum",
			Category: "spirits",
		},
		model.Ingredient{
			Quantity: 2,
			Name:     "puff pastry",
		},
	},
	Instructions: []model.Instruction{
		model.Instruction{"Let the butter soften at room temperature"},
		model.Instruction{"Mix the butter, flours, sugar, whole eggs and crème fraiche together in a bowl"},
		model.Instruction{"Put it on a first layer of puff pastry (keep an outer ring of 1 inch clear of the mixture"},
		model.Instruction{"Add the \"fève\"!"},
		model.Instruction{"Lay the second layer of puff pastry on top"},
		model.Instruction{"Press on the outer ring, so that the inside is airtight"},
		model.Instruction{"With a brush, put the egg yolk on top"},
		model.Instruction{"Bake in the oven at 210C during 20-25min"},
		model.Instruction{"While it's cooking, mix the water and confectioners' sugar"},
		model.Instruction{"When the galette is baked, spread the water/sugar mix on top with a brush"},
		model.Instruction{"Put it back in the oven for one more minute"},
	},
}
