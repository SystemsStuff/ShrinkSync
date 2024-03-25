import os, random, sys, logging

# generates a random word from the list of possible characters
def generate_word():
    possible_chars = 'abcdefghijklmnopqrstuvwxyz'
    word_length = random.randint(3, 5)
    return ''.join(random.choices(possible_chars, k = word_length))

"""
takes a list of elements of any type(strings/integers) and 
creates a new file and writes each element in a new line 
and raises an exception if the file already exists
"""
def write_to_file(file_name, data):
    if os.path.exists(file_name):
        raise Exception("Oops! The file: %s already exists :'(" % file_name)

    logging.info("Creating input file: %s" % file_name)
    with open(file_name, 'w') as f:
        f.write('\n'.join(data))

# generates all the necessary files given file
def generate_test_data(output_dir, file_prefix, file_count, line_count):
    partition_files = []
    for i in range(file_count):
        words = [generate_word() for _ in range(line_count)]
        partition_files.append(file_prefix + "_" + str(i))
        write_to_file(output_dir + "/" + partition_files[i], words)
    # creates the metadata file
    metadata_file = output_dir + "/" + file_prefix + "_metadata"
    logging.info("Creating metadata file: %s" % metadata_file)
    write_to_file(metadata_file, partition_files)

def main():
    output_dir, file_prefix, file_count, line_count = sys.argv[1], sys.argv[2], int(sys.argv[3]), int(sys.argv[4])
    if not os.path.exists(output_dir):
        os.makedirs(output_dir)
        logging.info("New Input directory: %s is created!" % output_dir)
    generate_test_data(output_dir, file_prefix, file_count, line_count)

if __name__ == "__main__":
    main()

