import time

def pascal_row(n):
    row = [1] * (n + 1)
    for k in range(1, n):
        row[k] = row[k - 1] * (n - k + 1) // k
    return row

def evaluate_polynomial(coeffs, x):
    result = 0
    steps = []
    for i, c in enumerate(coeffs):
        term = c * (x ** i)
        steps.append(f"{c} * {x}^{i} = {term}")
        result += term
    return result, steps

def main():
    n = 100
    x = 2
    start = time.time()
    coeffs = pascal_row(n)
    result, steps = evaluate_polynomial(coeffs, x)
    end = time.time()

    with open("Resultado2python.txt", "w") as f:
        f.write(f"Coeficientes de (x+1)^{n}:\n{coeffs}\n\n")
        f.write(f"Evaluación f({x}) = (x+1)^{n}:\n")
        for step in steps:
            f.write(step + "\n")
        f.write(f"\nResultado final: {result}\n")
        f.write(f"Tiempo de ejecución: {end - start:.6f} segundos\n")

if __name__ == "__main__":
    main()