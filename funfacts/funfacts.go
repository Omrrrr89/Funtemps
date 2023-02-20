package funfacts

// Sett inn alle Funfucts i en struktur
type FunFacts struct {
    Sun   []string
    Luna  []string
    Terra []string
    Facts []string
}

// Implementer funfacts-funksjon: GetFunFacts
func GetFunFacts(about string) []string {
  funFacts := FunFacts{
    Sun: []string{
      "The sun makes up 99.86% of the mass in the solar system.",
      "The sun is actually white, not yellow.",
      "It takes 8 minutes and 20 seconds for light to travel from the sun to Earth.",
    },
  
    Luna : []string{
      "The moon is not a perfect sphere, but is slightly flattened at the poles and bulges at the equator.",
      "The moon's gravity is about one-sixth that of Earth's gravity.",
      "The moon is moving away from the Earth at a rate of about 1.5 inches per year.",
    },
  
    Terra : []string{
      "Earth is the only planet known to support life.",
      "70% of the Earth's surface is covered in water.",
      "The tallest mountain on Earth, Mount Everest, is 29,029 feet tall.",
    },
 
    Facts : []string{
      "Sorry, I don't have any fun facts about that.",
    },
  
  }

switch about {

case "luna":
    return funFacts.Luna
case "terra":
  return funFacts.Terra

case "sun":
  return funFacts.Sun

default:
 return funFacts.Facts
}
}
