import SectionHeader from "@components/headers/sectionHeader/SectionHeader";
import ButtonSolid from "@components/buttons/Button/Solid";
import ButtonLink from "@components/buttons/Link";
import SearchTeachers from "./components/Search";

const SearchPage = (): JSX.Element => {
    return (
        <div>
            { /* Action Buttons */}

            <SectionHeader title="Selecione um professor">
                <>
                    <ButtonSolid title="Filtrar" sx={{ margin: "0 0 0 auto" }}/>
                    <ButtonLink
                        title="Criar professor"
                        variant="solid"
                        href="/admin/students/create"
                    />
                </>
            </SectionHeader>

            { /* Filter Student Form */}
            <SearchTeachers />
        </div>
    );
};

export default SearchPage;