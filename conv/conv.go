package conv

func FahrenheitToCelsius(fahrenheit float64) float64 {

  celsius := (fahrenheit - 32) * 5/9


    
    return celsius
}

func FahrenheitToKelvin(fahrenheit float64) float64 {

  Kelvin := (fahrenheit - 32) * 5/9 + 273.15
 
    return Kelvin
}

func CelsiusToFahrenheit(Celsius float64) float64 {

  fahrenheit := Celsius* 9/5 + 32
 
    return fahrenheit
}

func CelsiusToKelvin(Celsius float64) float64 {
  Kelvin := Celsius + 273.15

    return Kelvin
}

func KelvinToCelsius(Kelvin float64) float64 {
  Celsius := Kelvin - 273.15
 
    return Celsius
}

func KelvinToFahrenheit(Kelvin float64) float64 {
  fahrenheit := (Kelvin - 273.15)* 9/5 + 32

    return  fahrenheit

}