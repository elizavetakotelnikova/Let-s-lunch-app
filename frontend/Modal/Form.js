import React, {useState} from 'react'
import Button from "../Button/Button";
import Input from "../components/Input";

const Form =  () => {
    const [post, setPost] = useState({title: '', body: ''})
    const [title, setTitle] = useState('')
    const [posts, setPosts] = useState([])
    const [modal, setModal] = useState(false);

    const addNewPost = (e) => {
        e.preventDefault()
        const newPost = {
            ...post, id: Date.now()
        }
        setPost({title: '', body: ''})
    }

    const createPost = (newPost) => {
        setPosts([...posts, newPost])
        setModal(false)
    }

    return (
        <form>
            <Input
                value={post.title}
                onChange={e => setPost({...post, title: e.target.value})}
                type="text"
                placeholder="Ваше имя"
                />
            <Input
                value={post.body}
                onChange={e => setPost({...post, body: e.target.value})}
                type="text"
                placeholder="Время"
            />
            <Button onClick={addNewPost}>Создать встречу</Button>
        </form>
    );
};

export default Form;