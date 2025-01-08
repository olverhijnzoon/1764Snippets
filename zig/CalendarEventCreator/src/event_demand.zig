const c = @cImport({
    @cInclude("scheduler/event_demand.h");
});

pub export fn createEventDemand(summary: [*c]const u8, description: [*c]const u8, durationMinutes: i32, start: [*c]const u8, end: [*c]const u8) c.EventDemand {
    return c.EventDemand{
        .summary = summary,
        .description = description,
        .durationMinutes = durationMinutes,
        .timeSlot = c.TimeSlot{ .start = start, .end = end },
    };
}
