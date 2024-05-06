-- ignore

CREATE TABLE pessoa (
    ID_pessoa INT PRIMARY KEY AUTO_INCREMENT,
    Nome VARCHAR(100) NOT NULL ,
    Email VARCHAR(100) NOT NULL ,
    Telefone VARCHAR(20) NOT NULL 
);

CREATE TABLE livro (
    ID_livro INT PRIMARY KEY AUTO_INCREMENT,
    Titulo VARCHAR(100) NOT NULL ,
    Autor VARCHAR(100),
    AnoPublicacao INT 
);

CREATE TABLE livro (
    ID_reserva INT PRIMARY KEY AUTO_INCREMENT,
    DataReserva DATE NOT NULL ,
    ID_pessoa INT NOT NULL ,
    ID_livro INT NOT NULL ,
    FOREIGN KEY (ID_pessoa) REFERENCES pessoa(ID_pessoa),
    FOREIGN KEY (ID_livro) REFERENCES livro(ID_livro)
);
