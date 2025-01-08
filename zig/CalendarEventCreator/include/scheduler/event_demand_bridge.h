// Copyright 2025 Oliver Heinsohn

#ifndef INCLUDE_SCHEDULER_EVENT_DEMAND_BRIDGE_H_
#define INCLUDE_SCHEDULER_EVENT_DEMAND_BRIDGE_H_

#include "scheduler/event_demand.h"

extern EventDemand createEventDemand(const char *summary,
                                     const char *description,
                                     int durationMinutes, const char *start,
                                     const char *end);

#endif // INCLUDE_SCHEDULER_EVENT_DEMAND_BRIDGE_H_
