# README

## Projeto Ecossistema - Linguagens de Programação - 2019 / 02


![Ecossistema - Simulação de Competição de Recursos](https://github.com/luandkg/Ecossistema/tree/OrganismosXML/assets/prints/EcossistemaP3.png)


## Requisitos


 - Configure o GOPATH para a pasta "Ecossistema"

 - Motor Gráfico SDL : https://github.com/veandco/go-sdl2 
    - TERMINAL : apt install libsdl2{,-image,-mixer,-ttf,-gfx}-dev
    - TERMINAL: go get -v github.com/veandco/go-sdl2/{sdl,img,mix,ttf}
 
##Packages

    - BioXML
    - Ecossistema
    - Grafico
    - Tabuleiro
    - Utils

## Conceito Geral

Simulação de um ecossistema terrestre mostrando a interação de produtores e consumidores, entre si e com fatores abióticos, como chuva, vento, umidade e etc.
A simulação também implementa funções de ciclo de vida, como nascimento,movimentação, alimentação, reprodução e morte.

## Funções Implementadas

     - Carregamento automático de espécies
     - Geração randômica de organismos a partir das espécies
     - Mapeamento do espaço em Tabuleiro [50 x 50]
     - Interface Gráfica SDL
     - Interação de FATORES ABIÓTICOS
            - Tempo [ Ciclo - Dia e Noite ]
            - Vento
            - Umidade
            - Chuva
            - Temperatura
            - Nuvem
            - Luminosidade
            - Sensação Térmica
            
      - Ciclo de vida
             - Nascimento
             - Movimentação
             - Alimentação
             - Reprodução
             - Morte
                  
      - Interação entre Organismos
             - Predação
             
      - Plotagem Gráfica de FATORES ABIÓTICOS
      - Monitaração de Status
      - Sistema de logs
      - Pasta de Recursos [ Assets ]
            - Fonts
            - Organismos
       
## Funções Futuras

        - Relação entre organismos e FATORES ABIÓTICOS
        - População
        - Estruturação de dependência energética
              - Energia
              - Gás Oxigênio
              - Gás Carbônico
        - Plotagem Gráfica de Predatismo em Gráfico de Cruzamento Linear
        
### Gerando arquivo compilado

    - go build main
    - ./evolucao
