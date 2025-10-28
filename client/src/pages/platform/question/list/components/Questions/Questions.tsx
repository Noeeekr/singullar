import Filters from "./Filters";
import Search from "./Search";

export default function({ onParentClick }: { onParentClick: (cb: () => void) => void }): JSX.Element {
    return(
        <Filters onParentClick={onParentClick}>
            <Search />
        </Filters>
    )
}