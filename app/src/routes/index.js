import { Home, Login, Rank, Typing, User } from "../pages"

export const publicRoute = [
    {path: "/", element: Home},
    {path: "/home", element: Home},
    {path: "/login", element: Login},
    {path: "/rank", element: Rank},
    {path: "/typing", element: Typing},
]

export const privateRoute = [
    {path: "/user", element: User},
]