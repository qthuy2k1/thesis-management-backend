// Original file: src/proto/schedule.proto

import type { Long } from '@grpc/proto-loader';

export interface ReportingStageExerciseResponse {
  'id'?: (number | string | Long);
  'label'?: (string);
  'description'?: (string);
  'value'?: (string);
}

export interface ReportingStageExerciseResponse__Output {
  'id'?: (Long);
  'label'?: (string);
  'description'?: (string);
  'value'?: (string);
}
