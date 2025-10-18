import Button from "@components/buttons/Button";

const QuestionsPage = ():JSX.Element => {
    return(
        <div>
            Questions page
            <Button variant="button" type="button" title="Botão padrão (no variant)" icon={<div>Icone</div>}></Button>
            <Button variant="paper" type="button" title="Botão padrão (paper variant)" icon={<div>Icone</div>}></Button>
            
        </div>
    )
}
export default QuestionsPage;