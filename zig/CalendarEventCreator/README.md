# CalendarEventCreator

## Problem Statement

A calendar is a system for organizing events, appointments, and tasks, typically providing features like scheduling, reminders, and sharing. Consumer devices usually include a calendar application with a graphical user interface. While these interfaces offer features to reduce the effort involved in creating events, such as shortcuts, recurring events, and templates, they often lack the ability to dynamically adjust or synchronize with rapidly changing, non-static environments.

## Design Decisions

* Calendar events are created programmatically (through code).
* The ICS (iCalendar) standard file format (defined by RFC 5545) for representing calendar and scheduling information is used as it enables different calendar applications to exchange data.