// Copyright 2025 Oliver Heinsohn

#ifndef INCLUDE_SCHEDULER_EVENT_DEMAND_H_
#define INCLUDE_SCHEDULER_EVENT_DEMAND_H_

typedef struct {
  const char *start;
  const char *end;
} TimeSlot;

typedef struct {
  const char *summary;
  const char *description;
  int durationMinutes;
  TimeSlot timeSlot;
} EventDemand;

#endif // INCLUDE_SCHEDULER_EVENT_DEMAND_H_
