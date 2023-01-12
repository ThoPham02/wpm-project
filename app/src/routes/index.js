import Home from "../pages/Home"
import Login from "../pages/Login"
import Rank from "../pages/Rank"
import Typing from "../pages/Typing"
import User from "../pages/User"

export const publicRoute = [
    {path: "/", element: Home},
    {path: "/home", element: Home},
    {path: "/login", element: Login},
    {path: "/rank", element: Rank},
    {path: "/typing/:id", element: Typing},
]

export const privateRoute = [
    {path: "/user", element: User},
]