import fire
import uvicorn


def serve() -> None:
    uvicorn.run(
        'purse.main:app',
        host='0.0.0.0',
        port=80,
        reload=True,
    )


def main():
    fire.Fire(
        {
            'serve': serve,
        },
    )


if __name__ == '__main__':
    main()
