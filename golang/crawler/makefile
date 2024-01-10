# Number of HTML files to create
NUM_FILES = 10

# Directory to store HTML files
DIR = website

# List of HTML files
FILES = $(patsubst %,$(DIR)/file%.html,$(shell seq 1 $(NUM_FILES)))

all: $(DIR) $(FILES)

$(DIR):
	mkdir -p $(DIR)

$(DIR)/file%.html:
	echo '<!DOCTYPE html>' > $@
	echo '<html>' >> $@
	echo '<head><title>File $*</title></head>' >> $@
	echo '<body>' >> $@
	echo '<h1>File $*</h1>' >> $@
	$(foreach file,$(filter-out $@,$(FILES)),echo '<a href="$(notdir $(file))">$(notdir $(file))</a><br>' >> $@;)
	echo '</body>' >> $@
	echo '</html>' >> $@

clean:
	rm -rf $(DIR)

CRAWLER = golangcrawler
WEBSERVER = simplewebserver

build_go:
	go build -o $(CRAWLER)X $(CRAWLER).go
	go build -o $(WEBSERVER)/$(WEBSERVER)X $(WEBSERVER)/$(WEBSERVER).go

run_crawler_race:
	go run -race $(CRAWLER).go

run_crawler:
	./$(CRAWLER)X 

clean_after_crawler:
	rm -rf download