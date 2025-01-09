#include "raylib.h"

int main(void)
{
		const int screenWidth = 1764-420;
		const int screenHeight = 420;

		SetConfigFlags(FLAG_WINDOW_TRANSPARENT);
		InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window");
		SetTargetFPS(60);
		Color transparentBlack = {0, 0, 0, 69};
		while (!WindowShouldClose())
		{
			BeginDrawing();
				ClearBackground(transparentBlack);
				DrawText("1764Snippets", 400, 176, 42, GREEN);
			EndDrawing();
		}
		CloseWindow();
		return 0;
}
