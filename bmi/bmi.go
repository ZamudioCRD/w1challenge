package bmi

func CalcBmi(peso, altura float32) float32 {
	return peso / (altura * altura)
}

func BmiMensaje(imc float32) string {
    if imc < 18.5 {
        return "You are underweight, add more potato to the broth"
    } else if imc >= 18.5 && imc < 25 {
        return "You have a normal weight, I have healthy envy of you"
    } else {
        return "You are overweight, I know, the pandemic has affected us all"
    }
}