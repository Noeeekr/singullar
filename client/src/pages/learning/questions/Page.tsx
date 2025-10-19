import DefaultButton from "@components/buttons/Button/Default";

const QuestionsPage = ():JSX.Element => {
    return(
        <div>
            Questions page
            <DefaultButton title="Botão padrão (no variant)" icon={{ component: <div>Icone</div>}}></DefaultButton>
            <DefaultButton title="Botão padrão (paper variant)" icon={{ component: <div>Icone</div>}}></DefaultButton>
            
        </div>
    )
}
export default QuestionsPage;