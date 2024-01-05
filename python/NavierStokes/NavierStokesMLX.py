import mlx.core
import matplotlib.pyplot as plt
import time

def navier_stokes_2d(u, v, p, rho, nu, dt, dx, dy, nt):
    un = mlx.core.zeros_like(u)
    vn = mlx.core.zeros_like(v)

    calc_rate = 0
    start_time = time.time()
    
    for n in range(nt):
        un = mlx.core.array(u)
        vn = mlx.core.array(v)

        u[1:-1, 1:-1] = (un[1:-1, 1:-1] -
                        un[1:-1, 1:-1] * dt / dx *
                        (un[1:-1, 1:-1] - un[1:-1, 0:-2]) -
                        vn[1:-1, 1:-1] * dt / dy *
                        (un[1:-1, 1:-1] - un[0:-2, 1:-1]) -
                        dt / (2 * rho * dx) * (p[1:-1, 2:] - p[1:-1, 0:-2]) +
                        nu * (dt / dx**2 *
                            (un[1:-1, 2:] - 2 * un[1:-1, 1:-1] + un[1:-1, 0:-2]) +
                            dt / dy**2 *
                            (un[2:, 1:-1] - 2 * un[1:-1, 1:-1] + un[0:-2, 1:-1])))

        v[1:-1, 1:-1] = (vn[1:-1, 1:-1] -
                        un[1:-1, 1:-1] * dt / dx *
                        (vn[1:-1, 1:-1] - vn[1:-1, 0:-2]) -
                        vn[1:-1, 1:-1] * dt / dy *
                        (vn[1:-1, 1:-1] - vn[0:-2, 1:-1]) -
                        dt / (2 * rho * dy) * (p[2:, 1:-1] - p[0:-2, 1:-1]) +
                        nu * (dt / dx**2 *
                            (vn[1:-1, 2:] - 2 * vn[1:-1, 1:-1] + vn[1:-1, 0:-2]) +
                            dt / dy**2 *
                            (vn[2:, 1:-1] - 2 * vn[1:-1, 1:-1] + vn[0:-2, 1:-1])))

        p[1:-1, 1:-1] = (((un[1:-1, 2:] - un[1:-1, 0:-2]) *
                          (dt / (2 * dx))**2 +
                          (vn[2:, 1:-1] - vn[0:-2, 1:-1]) *
                          (dt / (2 * dy))**2) /
                         (2 * ((dt / dx)**2 + (dt / dy)**2)) -
                         ((un[1:-1, 2:] - un[1:-1, 0:-2]) *
                          dt / (2 * dx) +
                          (vn[2:, 1:-1] - vn[0:-2, 1:-1]) *
                          dt / (2 * dy)) * rho)

        # Boundary conditions
        u[0, :] = 0
        u[:, 0] = 0
        u[-1, :] = 0
        u[:, -1] = 0
        v[0, :] = 0
        v[:, 0] = 0
        v[-1, :] = 0
        v[:, -1] = 0
        u[41, 0] = 1
        u[42, 0] = 10
        u[43, 0] = 3
        v[0, 41] = -1
        v[0, 42] = -10
        v[0, 43] = -3

        if n % 1000 == 999:  # Every 1000 iterations
            end_time = time.time()
            calc_rate = 1000.0 / (end_time - start_time)
            start_time = time.time()

            # Plot the velocity field
            plt.clf()
            plt.quiver(u, v)
            plt.title(f'MLX, Iteration: {n+1}, Calculation Rate: {calc_rate:.2f} iterations/sec')
            plt.draw()
            plt.pause(0.01)

    return u, v, p

def main():
    # Define the grid
    nx = 50
    ny = 50
    dx = 2 / (nx - 1)
    dy = 2 / (ny - 1)

    # Define the physical parameters
    rho = 10
    nu = 0.1

    # Define the time step and number of time steps
    dt = 0.001
    nt = 100000

    # Initialize the velocity and pressure fields
    u = mlx.core.random.uniform(shape=(ny, nx))/100
    v = mlx.core.random.uniform(shape=(ny, nx))/100
    p = mlx.core.zeros(shape=(ny, nx))

    # Run the simulation
    u, v, p = navier_stokes_2d(u, v, p, rho, nu, dt, dx, dy, nt)

if __name__ == "__main__":
    main()
