import SearchTeachers from "./components/Search";
import SectionHeader from "@components/headers/sectionHeader/SectionHeader";
import ButtonSolid from "@components/buttons/Default/Solid";
import ButtonLink from "@components/buttons/Link";
import Stack from "@mui/material/Stack";

const SearchPage = (): JSX.Element => {
    return (
        <div>
            { /* Action Buttons */}
            <SectionHeader title="Selecione um professor">
                <Stack direction="row">
                    <ButtonSolid title="Filtrar" sx={{ margin: "0 0 0 auto" }} />
                    <ButtonLink
                        title="Criar professor"
                        variant="solid"
                        href="/admin/students/create"
                    />
                </Stack>
            </SectionHeader>

            { /* Filter Student Form */}
            <SearchTeachers />
        </div>
    );
};

export default SearchPage;