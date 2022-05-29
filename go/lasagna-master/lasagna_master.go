package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
  if time == 0 {
    time = 2
  }
  return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
  for _, layer := range layers {
    switch layer {
    case "noodles": noodles += 50
    case "sauce": sauce += 0.2
    }
  }
  return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(ingredients []string, withSecret []string) {
  secretIngredient := ingredients[len(ingredients) - 1]
  withSecret[len(withSecret) - 1] = secretIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
  scaled := make([]float64, len(quantities))
  for i := range scaled {
    scaled[i] = quantities[i] * (float64(portions) / 2)

  }
  return scaled
}
