import fire
import uvicorn


def serve() -> None:
    uvicorn.run(
        'purse.main:app',
        reload=True,
        port=8003,
    )


def main():
    fire.Fire(
        {
            'serve': serve,
        },
    )


if __name__ == '__main__':
    main()
