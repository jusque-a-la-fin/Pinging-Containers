import 'bootstrap/dist/css/bootstrap.min.css';
import Container from './Single';
import Containers from './Many';

const Home = () => {
  return (
    <div className="container mt-5 d-flex flex-column align-items-center">
       <h1 className="mb-4">Мониторинг Docker контейнеров</h1>
      <Containers/>
      <Container/>
    </div>
  );
};

export default Home;