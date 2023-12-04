from collections import Counter

def get_character_frequency(file_path):
    try:
        # Read the content of the file
        with open(file_path, 'r', encoding='utf-8') as file:
            content = file.read()

        # Use Counter to count the frequency of each character
        character_frequency = Counter(content)

        # Print the results
        for char, freq in character_frequency.items():
            print(f"Character: '{char}', Frequency: {freq}")

        return character_frequency

    except FileNotFoundError:
        print(f"File not found: {file_path}")
        return None

# Replace 'your_file.txt' with the path to your text file
file_path = 'input.txt'
get_character_frequency(file_path)

