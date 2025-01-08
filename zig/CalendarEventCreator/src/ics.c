// Copyright 2025 Oliver Heinsohn

#include "ics/ics.h"
#include <stdlib.h>
#include <string.h>

Event createEvent(const char *dtstart, const char *dtend, const char *summary,
                  const char *description) {
  Event newEvent;
  strncpy(newEvent.dtstart, dtstart, sizeof(newEvent.dtstart));
  strncpy(newEvent.dtend, dtend, sizeof(newEvent.dtend));
  strncpy(newEvent.summary, summary, sizeof(newEvent.summary));
  strncpy(newEvent.description, description, sizeof(newEvent.description));
  return newEvent;
}

int createICSFile(Event event) {
  char *outputDir = getenv("ICS_OUTPUT_DIR");
  if (outputDir == NULL) {
    fprintf(stderr, "[ERROR] ICS_OUTPUT_DIR environment variable not set.\n");
    return 1;
  }
  printf("[INFO] Environment variable ICS_OUTPUT_DIR is set to '%s'.\n",
         outputDir);

  char filepath[256];
  snprintf(filepath, sizeof(filepath), "%s/event.ics", outputDir);
  FILE *file = fopen(filepath, "w");
  if (file == NULL) {
    fprintf(stderr, "[ERROR] Could not open file '%s' for writing.\n",
            filepath);
    return 1;
  }

  fprintf(file,
          "BEGIN:VCALENDAR\n"
          "VERSION:2.0\n"
          "BEGIN:VEVENT\n"
          "DTSTART:%s\n"
          "DTEND:%s\n"
          "SUMMARY:%s\n"
          "DESCRIPTION:%s\n"
          "END:VEVENT\n"
          "END:VCALENDAR\n",
          event.dtstart, event.dtend, event.summary, event.description);

  fclose(file);
  printf("[INFO] ICS file created successfully at '%s'.\n", filepath);
  return 0;
}

int writeICSFile() { return 0; }
