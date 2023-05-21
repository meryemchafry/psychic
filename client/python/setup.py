from setuptools import setup

# read the contents of your README file
from os import path
this_directory = path.abspath(path.dirname(__file__))
with open(path.join(this_directory, 'README.md'), encoding='utf-8') as f:
    long_description = f.read()

setup(
    name='psychicapi',
    version='0.4',
    description='Psychic.dev is an open-source universal data connector for knowledgebases.',
    long_description=long_description,
    long_description_content_type='text/markdown',
    author='Ayan Bandyopadhyay',
    author_email='ayan@getsidekick.ai',
    packages=['psychicapi'],
    install_requires=[
        'requests',
    ],
)