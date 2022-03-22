import './Button.styles.css';
const Button = ({placeholder, color}) => {
    return (
        <button className={"btn " + color}>{placeholder}</button>
    )
}

export default Button;