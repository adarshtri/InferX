#ifndef BRIDGE_H
#define BRIDGE_H

#ifdef __cplusplus
extern "C" {
#endif

// ProcessBatch simulates an inference batch processing.
// Returns 0 on success.
int ProcessBatch(int batch_size);

#ifdef __cplusplus
}
#endif

#endif // BRIDGE_H
