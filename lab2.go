package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unsafe"
)

func format_input() ([]byte, int) {

	var result [100]byte

	//Ask for input and separate it in bytes
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string (Max 100 characters): ")
	raw_input, _ := reader.ReadString('\n')
	input_bytes := []byte(raw_input)

	//Inicialize pointers
	ptr_input := &input_bytes[0]
	ptr_result := &result[0]

	count := 0
	valid_count := 0

	for {
		//Character as the value of input pointer
		character := *ptr_input

		if character == '\n' || count >= 100 {
			break
		}

		if (character >= 'A' && character <= 'Z') || (character >= 'a' && character <= 'z') || (character == ' ' && valid_count > 0) {
			//Value of result pointer change to character
			*ptr_result = character
			//Step up result pointer
			ptr_result = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr_result)) + 1))
			valid_count++
		}

		count++

		//Step up input pointer
		ptr_input = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr_input)) + 1))
	}

	return result[:valid_count], valid_count

}

func reverse_list(letter_list *[]byte, length int) {

	if length <= 1 {
		return
	}

	//Pointers at start and end of the string
	left_ptr := &(*letter_list)[0]
	right_ptr := (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(left_ptr)) + uintptr(length-1)))

	count := 0

	for count < length/2 {

		*left_ptr, *right_ptr = *right_ptr, *left_ptr

		//For doing pointer arithmetic
		left_ptr = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(left_ptr)) + 1))
		right_ptr = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(right_ptr)) - 1))

		count++
	}

}

func count_vowels_consonants(letter_list *[]byte, length int) (map[byte]int, int, int) {

	vowels_counts := map[byte]int{}
	vowels_bytes := [5]byte{'A', 'E', 'I', 'O', 'U'}
	number_of_vowels := 0
	number_of_consonants := 0

	ptr_letter := &(*letter_list)[0]

	count_letter_list := 0

	//For to traverse letter list
	for {
		character := (*ptr_letter)
		ptr_vowels := &vowels_bytes[0]

		if count_letter_list == length {
			break
		}

		count_vowels_list := 0
		is_vowel := false

		//For to traverse and compare vowels
		for {
			vowel := (*ptr_vowels)

			if count_vowels_list == 5 {
				break
			}

			//Compare vowels in lower and upper
			if character == vowel || character == vowel+32 {

				vowels_counts[vowel+32] += 1
				number_of_vowels++
				is_vowel = true

				break
			}

			ptr_vowels = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr_vowels)) + 1))
			count_vowels_list++
		}

		if character != ' ' && !is_vowel {
			number_of_consonants++
		}

		ptr_letter = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr_letter)) + 1))
		count_letter_list++

	}

	return vowels_counts, number_of_vowels, number_of_consonants

}

func format_vowels_consonants_output(vowels, consonants int, type_vowels string) string {

	spaced := strings.Join(strings.Split(type_vowels, ""), " ")
	output := fmt.Sprintf(" %d %s %d ", vowels, spaced, consonants)

	return output
}

func sort_map(vowels_map map[byte]int) string {
	vowels_order := [5]byte{'a', 'e', 'i', 'o', 'u'}

	ptr := &vowels_order[0]
	count := 0

	output := ""

	for {

		if count == 5 {
			break
		}

		vowel := *ptr
		value := vowels_map[vowel]
		if value > 0 {
			output += fmt.Sprint(value)
		}

		ptr = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + 1))
		count++

	}

	return output
}

func change_space_for_underscores(letter_list *[]byte, length int) {

	ptr_letter := &(*letter_list)[0]
	count := 0

	for {
		character := *ptr_letter

		if length == count {
			break
		}

		if character == ' ' {
			*ptr_letter = '_'
		}

		count++
		ptr_letter = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr_letter)) + 1))
	}

}

func main() {

	//Clean string input converting it in a slice
	input_list, length := format_input()

	reverse_list(&input_list, length)
	fmt.Print(string(input_list))

	vowels_counts, number_of_vowels, number_of_consonants := count_vowels_consonants(&input_list, length)
	ordered_vowels_count := sort_map(vowels_counts)
	fmt.Print(format_vowels_consonants_output(number_of_vowels, number_of_consonants, ordered_vowels_count))

	change_space_for_underscores(&input_list, length)
	fmt.Print(string(input_list))
}
