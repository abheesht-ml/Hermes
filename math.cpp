#include "vector_math.h"
#include <cmath>

extern "C" float euclidean_distance(float *a, float *b, int n) {
  float sum = 0.0f;
  for (int i = 0; i < n; i++) {
    float diff = a[i] - b[i];
    sum += diff * diff;
  }
  return std::sqrt(sum);
}