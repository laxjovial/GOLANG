package piscine

func TrimAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	var intstr string = ""
	for _, i := range s {
		if i == '-' && len(intstr) == 0 {
			intstr += "-"
		}
		if i < '0' || i > '9' {
			continue
		}
		intstr += string(i)
	}
	if len(intstr) == 0 || intstr == "-" {
		return 0
	}
	return Atoi(intstr)
}
