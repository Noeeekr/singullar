// Features
import { useState } from 'react';
import { useNavigate } from 'react-router-dom'
import useInstitutions from '../../../hooks/useInstitutions';

// Components
import Typography from '@mui/material/Typography';
import FormControl from '@mui/material/FormControl';
import InputLabel from '@mui/material/InputLabel';
import Select from '@mui/material/Select';
import MenuItem from '@mui/material/MenuItem';

const InstitutionSelection = (): JSX.Element => {
    const [institutions] = useInstitutions();
    const [selectValue, setSelectValue] = useState("");
    const navigate = useNavigate();

    return (
        <div>
            <Typography component="p">
                Você tem {institutions.length ? institutions.length : "?"} escolas em sua rede
            </Typography>
            <form>
                <FormControl
                    aria-labelledby="school-select-label"
                    sx={{
                        marginY: 1
                    }}
                >
                    <InputLabel id="school-select-label" sx={{
                        zIndex: 5,
                    }}>
                        Selecione uma escola
                    </InputLabel>
                    <Select
                        labelId="school-select-label"
                        id="school-select"
                        label="Selecione uma escola"
                        onChange={(e) => {
                            setSelectValue(e.target.value)

                            navigate("/admin/classes/" + e.target.value);
                        }}
                        value={selectValue}
                    >
                        {
                            institutions.map((inst) => {
                                return <MenuItem
                                    value={inst.id}
                                    key={inst.id + inst.name}
                                >{inst.name}</MenuItem>
                            })
                        }
                    </Select>
                </FormControl>
            </form>
        </div>
    )
}

export default InstitutionSelection