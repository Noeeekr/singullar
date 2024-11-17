// Features
import { useState, useEffect, useRef } from 'react';
import { useNavigate } from 'react-router-dom'
import useInstitutions from '../../../hooks/useInstitutions';

// Components
import Typography from '@mui/material/Typography';
import FormControl from '@mui/material/FormControl';
import InputLabel from '@mui/material/InputLabel';
import Select from '@mui/material/Select';
import MenuItem from '@mui/material/MenuItem';

const InstitutionSelection = (): JSX.Element => {
    const [institutions, getInstitutions] = useInstitutions();
    const hasFetched = useRef(false)
    
    const [selectValue, setSelectValue] = useState("");
    
    const navigate = useNavigate();

    useEffect(() => {
        if (!hasFetched.current) {
            hasFetched.current = true;
            getInstitutions()
        }
    },[])

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
                            
                            if (e.target.value) {
                                navigate("/admin/classes/" + e.target.value);
                            }
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
                        {
                            Boolean(institutions.length)
                                ? []
                                : <MenuItem
                                    value=""
                                    key="sometextthatisnotainstname"
                            >Procurando..</MenuItem>
                        }
                    </Select>
                </FormControl>
            </form>
        </div>
    )
}

export default InstitutionSelection