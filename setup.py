import setuptools

setuptools.setup(
    name='my_purse',
    version='0.0.1',
    author='Honchar',
    description=(
        'Store bank transactions into purse.'
    ),
    url='https://github.com/honcharevskiy/my_purse',
    packages=setuptools.find_packages(
        exclude=[
            'tests',
            'tests.*',
        ],
    ),
    install_requires=[
        # API
        'mangum',
        'fastapi',
        'pydantic',
        'uvicorn',

        # CLI
        'fire',

        # Monitoring
        'sentry_sdk',

    ],
    extras_require={
        'dev': [
            'pytest',
            'requests',
        ]
    },

    classifiers=[
        'Programming Language :: Python :: 3',
        'License :: Other/Proprietary License',
        'Operating System :: OS Independent',
    ],
)
