//go get github.com/brianvoe/gofakeit/v6

package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"reflect"
	"math/rand"
	"strconv"
)

type Usuario struct {
	cpf string
	email string
	senha string 
} 

type Professor struct {
	nome string
	cpf string
} 

type Aluno struct {
	nome string
	cpf string
} 

type Curso struct {
	nome string
	cpf_autor string
	id string 
} 

type Certificado struct {
	id string
	horas string
	} 

type Aluno_curso struct {
	cpf_aluno string
	id_curso string
	data_in string
	data_fim string 
	id_certificado string
}

func randInt(min, max int) int {
    return min + rand.Intn(max-min)
}

func valorExiste(valor string, lista interface{}, campo string) bool {
	v := reflect.ValueOf(lista) // retorna um objeto reflect, permitindo inspecionar o tipo, os campos e até o conteúdo de forma dinâmica

	for i := 0; i < v.Len(); i++ {
		item := v.Index(i) // retorna o i elemento da slice, ex: um item do struct, mas ainda sendo um reflect.value
		// Acessar o campo
		c := item.FieldByName(campo) // campo que estou comparando (ex: cpf)
		// Comparar se é string
		if c.String() == valor { 
			return true
		}
	}
	return false
}//Retorna verdadeiro se o valor ja existir no slice (vulgo lista)
//fmt.Println(valorExiste("123", usuarios, "cpf"))

func numExiste(num int, lista []int) bool{
	for i := 0; i < len(lista); i++{
		if num == lista[i]{
			return true
		}
	}
	return false
}

func stringExiste(palavra string, lista []string) bool{
	for i := 0; i < len(lista); i++{
		if palavra == lista[i]{
			return true
		}
	}
	return false
}

func gerarUsuarios(n int) []Usuario{
	usuarios := []Usuario{}
	for len(usuarios) < n {
		cpf := gofakeit.Regex("[0-9]{11}")
		email := gofakeit.Email()
		if valorExiste(cpf, usuarios, "cpf") || valorExiste(email, usuarios, "email"){
			continue // já existe → gera outro
		}
		u := Usuario{cpf: cpf, email: email, senha: gofakeit.Password(true, true, true, false, false, 6)}
		usuarios = append(usuarios, u)
	}
	return usuarios
}

func gerarAlunos(usuarios []Usuario, lista []int) []Aluno{
	alunos := []Aluno{}
	n_alunos := len(lista)
	for len(alunos) < n_alunos{
		nome := gofakeit.Name()
		if valorExiste(nome, alunos, "nome"){
			continue
		}
		index := lista[0]
		cpf := usuarios[index].cpf
		lista = lista[1:]
		a:= Aluno{nome: nome, cpf: cpf}
		alunos = append(alunos, a)
	}
	return alunos
}

func gerarProfessores(usuarios []Usuario, lista []int) []Professor{
	professores := []Professor{}
	n_professores := len(lista)
	for len(professores) < n_professores{
		nome := gofakeit.Name()
		if valorExiste(nome, professores, "nome"){
			continue
		}
		index := lista[0]
		cpf := usuarios[index].cpf
		lista = lista[1:]
		p:= Professor{nome: nome, cpf: cpf}
		professores = append(professores, p)
	}
	return professores
}

func gerarCursos(professores []Professor) []Curso{
	cursos := []Curso{}
	n_cursos := len(professores) + (len(professores)/2)
	cpfs := []string{}
	for i:= 0; i < len(professores); i++{
		cpf := professores[i].cpf
		cpfs = append(cpfs, cpf)
	}
	lista := make([]int, len(cpfs)) // lista que vai guardar a quant de cursos de cada professor
	nomes := []string{}
	adjectives := []string{"Introdução à ", " Avançado"}  
	for i := 0; i < n_cursos; i++{
		id := "CU-" + strconv.Itoa(i + 1) 

		cpf := ""
		index := 0
		aux :=0
		for aux == 0{
			if len(cpfs) != 1{
				index = randInt(0, len(cpfs))
			}
			lista[index] += 1;
			if lista[index] <= 2{
				cpf = cpfs[index]
				aux = 1
			}  
		}
		aux1 := 0
		nome := ""
		for aux1 == 0{
			nome = gofakeit.ProgrammingLanguage()
			if stringExiste(nome, nomes){
				continue
			}
			nomes = append(nomes, nome)
			aux1 = 1
		}
		n := randInt(0,2)
		if n == 0{
			nome = adjectives[0] + nome
		}else{
			nome = nome + adjectives[1]
		} 

		c := Curso{nome: nome, cpf_autor: cpf, id: id}
		cursos = append(cursos, c)
	}
	return cursos

}

func main(){
	//minimo de usuarios é 4 e n sei por que 6 e 7 tbm não vai
	// dps fazer um randon para gerar o numero de usuarios (8 - 20)
	usuarios := gerarUsuarios(10)
	lista := []int{}
	tam_usuario := len(usuarios);
	for len(lista) < tam_usuario{
		num := randInt(0,tam_usuario)
		if numExiste(num, lista){
			continue
		}
		lista = append(lista, num)
	}
	lista1 := lista[:(len(lista)/2) - 1]
	lista2 := lista[(len(lista1)):]
	professores := gerarProfessores(usuarios, lista1)
	alunos := gerarAlunos(usuarios, lista2) 

	fmt.Println(professores)
	fmt.Println(alunos)

	cursos := gerarCursos(professores)
	fmt.Println(cursos)


	

	

}