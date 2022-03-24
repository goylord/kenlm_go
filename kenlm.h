void* load_model(const char *file_name);
float score(void* model, char *sentence, int bos, int eos);
float perplexity(void* model, char *sentence);