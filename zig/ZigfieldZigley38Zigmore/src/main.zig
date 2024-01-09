const std = @import("std");
const print = std.debug.print;

// game state
var gameEpoch: u8 = 0;
var finished: bool = !true;

// game config
const display_delay: u64 = 50; // milliseconds

const introtext = "Axel, as you step through the grand oak doors of our esteemed law firm, know that you're not just joining a company, but a legacy. The hallowed halls have witnessed countless victories, late-night strategy sessions, and the rise and fall of many a legal titan. Your desk, bathed in the soft glow of the chandelier, awaits your touch. The library, with its centuries-old legal tomes, beckons. The partners, Zigfield, Zigley, and Zigmore, have heard of your prowess and eagerly anticipate the mark you'll leave. So, grab your robe, power up your terminal, and let's script history together. Welcome to the team, Axel!\n\n\n";

pub fn main() !void {
    // Clear the terminal and set text color to green using ANSI escape sequences
    print("\x1B[H\x1B[2J\x1B[32m", .{});

    display_introtext();
    display_game_state();
}

pub fn display_introtext() void {
    for (introtext) |byte| {
        print("{c}", .{byte});
        std.time.sleep(display_delay * std.time.ns_per_ms);
    }
    gameEpoch += 1;
}

pub fn display_game_state() void {
    print("Epoch: {}\n", .{gameEpoch});
    print("Finished? {}\n", .{finished});
}