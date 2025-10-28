import Stack from "@mui/material/Stack"
import { styled } from "@mui/material"

export default styled(Stack)(() => ({
    alignItems: "center",
    justifyContent: "space-between",

    backgroundColor: "rgba(130, 130, 130, 0.1)",
    padding: 18,
    border: "solid 1px rgb(190,190,190)",
    borderBottom: "none",
    borderTopLeftRadius: 16,
    borderTopRightRadius: 16,
}))