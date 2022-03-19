void* load_model(const char *file_name);
float score(void* model, char *sentence);
float perplexity(void* model, char *sentence);