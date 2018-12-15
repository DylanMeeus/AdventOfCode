; 14 of AoC

(defparameter *loops* 10)

; get the digits of a number
(defun digits(n) 
  (map 'list #'digit-char-p (prin1-to-string n)))


(defun next-index(lst pos score)
  (mod (+ pos (1+ score)) (list-length lst)))

; generate recipes
(defun generate(input iteration fst snd)
  ; just test with adding one for starters
  (let* ((x (nth fst input))
        (y (nth snd input))
        (sm (digits (+ x y))))
  (setf input (append input sm))
  (if (= iteration *loops*)
    input 
    (generate input (1+ iteration) (next-index input fst x) (next-index input snd y))
  )))



