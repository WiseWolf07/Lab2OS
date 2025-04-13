#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

void format_input(char *result, int *valid_count) {
    char input[200];
    printf("Enter a string (Max 100 characters): ");
    fgets(input, sizeof(input), stdin);

    int count = 0;
    *valid_count = 0;

    for (int i = 0; i < 100 && input[i] != '\0' && input[i] != '\n'; i++) {
        char character = input[i];
        if (isalpha(character) || (character == ' ' && *valid_count > 0)) {
            result[*valid_count] = character;
            (*valid_count)++;
        }
        else{
            break;
        }
    }
}

void reverse_list(char *list, int length) {
    for (int i = 0; i < length / 2; i++) {
        char temp = list[i];
        list[i] = list[length - i - 1];
        list[length - i - 1] = temp;
    }
}

void count_vowels_consonants(char *list, int length, int *vowel_counts, int *num_vowels, int *num_consonants) {
    char vowels[5] = {'A', 'E', 'I', 'O', 'U'};
    *num_vowels = 0;
    *num_consonants = 0;
    memset(vowel_counts, 0, 5 * sizeof(int));

    for (int i = 0; i < length; i++) {
        char c = toupper(list[i]);
        int is_vowel = 0;
        for (int j = 0; j < 5; j++) {
            if (c == vowels[j]) {
                vowel_counts[j]++;
                (*num_vowels)++;
                is_vowel = 1;
                break;
            }
        }
        if (!is_vowel && c != ' ') {
            (*num_consonants)++;
        }
    }
}

void format_vowels_consonants_output(int vowels, int consonants, int *vowel_counts, char *output) {
    sprintf(output, " %d ", vowels);
    for (int i = 0; i < 5; i++) {
        if (vowel_counts[i] > 0) {
            char temp[10];
            sprintf(temp, "%d ", vowel_counts[i]);
            strcat(output, temp);
        }
    }
    char consonant_str[10];
    sprintf(consonant_str, "%d ", consonants);
    strcat(output, consonant_str);
}

void change_space_for_underscores(char *list, int length) {
    for (int i = 0; i < length; i++) {
        if (list[i] == ' ') {
            list[i] = '_';
        }
    }
}

int main() {
    char result[100] = {0};
    int valid_count;

    format_input(result, &valid_count);

    reverse_list(result, valid_count);
    printf("%.*s", valid_count, result);

    int vowel_counts[5];
    int num_vowels, num_consonants;
    count_vowels_consonants(result, valid_count, vowel_counts, &num_vowels, &num_consonants);

    char formatted_output[100] = {0};
    format_vowels_consonants_output(num_vowels, num_consonants, vowel_counts, formatted_output);
    printf("%s", formatted_output);

    change_space_for_underscores(result, valid_count);
    printf("%.*s", valid_count, result);

    return 0;
}
