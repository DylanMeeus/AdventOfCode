; first puzzle of AoC 2018

(defun get-file()
  (with-open-file (stream "input.txt")
    (loop for line = (read-line stream nil)
          while line
          collect line)))


(defun solve() 
  (let ((input (mapcar (lambda (s) (parse-integer s)) (get-file))))
    (apply '+ input)))
