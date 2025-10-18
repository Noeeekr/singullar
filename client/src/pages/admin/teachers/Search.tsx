import SectionHeader from "@components/headers/sectionHeader/SectionHeader";
import Button from "@components/buttons/ButtonSolid";
import ButtonLink from "@components/buttons/ButtonLink";
import SearchTeachers from "./components/Search";

const SearchPage = (): JSX.Element => {
    return (
        <div>
            { /* Action Buttons */}

            <SectionHeader title="Selecione um professor">
                <>
                    <Button sx={{ margin: "0 0 0 auto" }}>Filtrar</Button>
                    <ButtonLink
                        icon={<></>}
                        type="link"
                        title="doesnt-matter"
                        href="/admin/students/create"
                    >
                        Criar professor
                    </ButtonLink>
                </>
            </SectionHeader>

            { /* Filter Student Form */}
            <SearchTeachers />
        </div>
    );
};

export default SearchPage;