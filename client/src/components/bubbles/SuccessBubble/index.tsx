import ErrorBubble from "../ErrorBubble"

export default ( { message }: { message: string }) => {
    return <ErrorBubble message={message} primary="rgba(155, 226, 141, 1)" secondary="rgba(47, 105, 29, 1)"/>
}