#ifndef BRIDGE_H
#define BRIDGE_H

#ifdef __cplusplus
extern "C" {
#endif

// ProcessBatch processes an inference batch using real metadata.
// Returns 0 on success.
int ProcessBatch(const int* prompt_lengths, int batch_size);

#ifdef __cplusplus
}
#endif

#endif // BRIDGE_H
